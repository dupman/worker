/*
 * This file is part of the dupman/worker project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package task

import (
	"log"

	"github.com/dupman/celery"
	"github.com/gocelery/gocelery"
)

type WebsiteTask struct {
	celery.Task
}

func NewWebsiteTask(client *gocelery.CeleryClient) celery.TaskInterface {
	task := &WebsiteTask{
		celery.Task{
			Client:    client,
			Namespace: "website",
			Functions: map[string]interface{}{},
		},
	}

	task.Functions["fetch"] = task.fetch

	return task
}

func (t *WebsiteTask) fetch(id, url, token string) {
	// @todo: do work...
	log.Printf("[website] fetching %s", url)
}
