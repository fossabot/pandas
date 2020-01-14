//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package cache

import (
	"testing"

	"github.com/cloustone/pandas/models"
	modelsoptions "github.com/cloustone/pandas/models/options"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCacheSet(t *testing.T) {
	servingOptions := modelsoptions.NewServingOptions()
	servingOptions.Cache = modelsoptions.KCacheRedis
	cache := NewCache(servingOptions)
	viewID := "id12345"

	origin := &models.View{
		Name:      "hello",
		ProjectID: "project1",
	}

	Convey("TestCacheSet should return ok when two cache items are same", t, func() {
		err := cache.Set(viewID, origin)
		So(err, ShouldBeNil)

		view := models.View{}
		err = cache.Get(viewID, &view)
		So(err, ShouldBeNil)
		//	So(StringSliceEqual(origin.Name, view.Name), ShouldBeTrue)
		//	So(StringSliceEqual(origin.ProjectID, view.ProjectID), ShouldBeTrue)
	})

	Convey("TestCacheSet should return false when two cache items dismatched", t, func() {
		err := cache.Set(viewID, origin)
		So(err, ShouldBeNil)

		view := models.View{}
		err = cache.Get(viewID+"1", &view)
		So(err, ShouldNotBeNil)
	})

}

func TestCacheSetAndGet(t *testing.T) {
	servingOptions := modelsoptions.NewServingOptions()
	servingOptions.Cache = modelsoptions.KCacheRedis
	cache := NewCache(servingOptions)
	viewID := "id12345"
	origin := &models.View{
		Name:      "hello",
		ProjectID: "project1",
	}
	Convey("mofang testcache set return true when two cache items are same", t, func() {
		err := cache.Set(viewID, origin)
		So(err, ShouldBeNil)

		view := models.View{}
		err = cache.Get(viewID, &view)
		So(err, ShouldBeNil)
	})

	Convey("mofang testcache set&get return false when two cache items are dismatched", t, func() {
		err := cache.Set(viewID, origin)
		So(err, ShouldBeNil)

		view := models.View{}
		err = cache.Get(viewID+"1", &view)
		So(err, ShouldNotBeNil)
	})
}

func TestCacheDelete(t *testing.T) {
	servingOptions := modelsoptions.NewServingOptions()
	servingOptions.Cache = modelsoptions.KCacheRedis
	cache := NewCache(servingOptions)

	viewID := "id12345"
	origin := &models.View{
		Name:      "Hello",
		ProjectID: "project1",
	}
	Convey("test cache delete return true when cache item is deleted", t, func() {
		err := cache.Set(viewID, origin)
		So(err, ShouldBeNil)

		view := models.View{}
		err = cache.Get(viewID, &view)
		So(err, ShouldBeNil)

		err = cache.Delete(viewID)
		So(err, ShouldBeNil)

		view = models.View{}
		err = cache.Get(viewID, &view)
		So(err, ShouldNotBeNil)
	})
}

func TestListPush(t *testing.T) {
	servingOptions := modelsoptions.NewServingOptions()
	servingOptions.Cache = modelsoptions.KCacheRedis
	cache := NewCache(servingOptions)

	viewID := "id123456"

	origin1 := &models.View{
		Name:      "Hello",
		ProjectID: "project1",
	}

	origin2 := &models.View{
		Name:      "GoodBye",
		ProjectID: "project2",
	}

	Convey("test listpush", t, func() {
		err := cache.ListPush(viewID, origin1, origin2)
		So(err, ShouldNotBeNil)

		view := models.View{}
		err = cache.Get(viewID, &view)
		So(err, ShouldNotBeNil)
	})
}
