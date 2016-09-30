package githubapi

import (
	"io/ioutil"
	"net/http"

	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/go-errors/errors"
	"net/url"
	"strings"
)

const configListUrl string = "https://github.schq.secious.com/api/v3/repos/Logrhythm/ConfigElements/contents/src/main/java/com/logrhythm/configelements?ref=%s"

type LRGithubConfig struct {
	httpGet func(string) (*http.Response, error)
}

func NewLRGithubConfig() *LRGithubConfig {
	return &LRGithubConfig{
		httpGet: http.Get,
	}
}

func (c *LRGithubConfig) ConfigFilesForBranch(branch string) (map[string]*url.URL, error) {

	resp, err := http.Get(fmt.Sprintf(configListUrl, url.QueryEscape(branch)))
	if err != nil {
		return nil, errors.Errorf("listConfigFiles Get error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("listConfigFiles ReadAll error: %v", err)
	}
	//	log.Debugf("body: %s", string(body))

	contents := []ContentsResponse{}
	err = json.Unmarshal(body, &contents)
	if err != nil {
		return nil, errors.Errorf("listConfigFiles json Unmarshal error: %v", err)
	}

	configs := map[string]*url.URL{}
	for _, c := range contents {
		//		log.Debugf("%-12s %-24s %s", c.Type, c.Name, c.DownloadURL)
		if strings.Compare(c.Type, "file") == 0 && strings.HasSuffix(c.Name, ".yaml") {
			u, err := url.Parse(c.DownloadURL)
			if err != nil {
				return configs, errors.Errorf("listConfigFiles url Parse error for '%s': %v", c.DownloadURL, err)
			}
			configs[c.Name] = u
		}
	}

	for k, v := range configs {
		log.Debug(k, v)
	}

	return configs, nil
}
