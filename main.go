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

package main

import (
	"log"

	"github.com/dupman/celery"
	"github.com/dupman/worker/lib"
	"github.com/dupman/worker/task"
)

func main() {
	conf := lib.NewConfig()

	celeryClient, err := celery.NewClient(&conf.Celery)
	if err != nil {
		log.Fatalf("Failed to create celery client: %s", err)
	}

	log.Println("Registering tasks")
	celery.RegisterTasks(celeryClient, task.NewWebsiteTask)

	log.Println("Starting dupman worker")
	celeryClient.StartWorker()
	celeryClient.WaitForStopWorker()
}
