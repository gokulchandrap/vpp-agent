// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linuxplugin

import (
	"os"
	"strconv"
	"time"

	log "github.com/ligato/cn-infra/logging/logrus"
	"golang.org/x/net/context"
)

var (
	sleepAfterLinuxResync = time.Duration(0)
)

func init() {
	sec, err := strconv.Atoi(os.Getenv("SLEEP_AFTER_LINUX_RESYNC"))
	if err == nil && sec > 0 {
		sleepAfterLinuxResync = time.Second * time.Duration(sec)
	}
}

// WatchEvents goroutine is used to watch for changes in the northbound configuration
func (plugin *Plugin) watchEvents(ctx context.Context) {
	plugin.wg.Add(1)
	defer plugin.wg.Done()

	for {
		select {
		case resyncEv := <-plugin.resyncChan:
			req := resyncParseEvent(resyncEv)
			err := plugin.resyncPropageRequest(req)

			resyncEv.Done(err)

			// optional hard sleep after linux resync
			if sleepAfterLinuxResync > 0 {
				log.DefaultLogger().Warnf("starting sleep after linux resync for %v", sleepAfterLinuxResync)
				time.Sleep(sleepAfterLinuxResync)
				log.DefaultLogger().Warnf("finished sleep after linux resync")
			}

		case dataChng := <-plugin.changeChan:
			err := plugin.changePropagateRequest(dataChng)

			dataChng.Done(err)

		case <-ctx.Done():
			log.DefaultLogger().Debug("Stop watching events")
			return
		}
	}
}
