package model_test

import (
	. "github.com/michaellihs/golab/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/json"
)

var _ = Describe("Project", func() {

	It("can be unmarshaled from JSON with project_acces", func() {
		project := new(Project)
		json.Unmarshal([]byte(`{
		    "id": 2859355,
		    "description": null,
		    "default_branch": null,
		    "tag_list": [],
		    "public": false,
		    "archived": false,
		    "visibility_level": 0,
		    "ssh_url_to_repo": "git@gitlab.com:michaellihs/test.git",
		    "http_url_to_repo": "https://gitlab.com/michaellihs/test.git",
		    "web_url": "https://gitlab.com/michaellihs/test",
		    "owner": {
		      "name": "Michael Lihs",
		      "username": "michaellihs",
		      "id": 1150266,
		      "state": "active",
		      "avatar_url": "https://secure.gravatar.com/avatar/4e1d79b0859761cf6de8f560784226cc?s=80\u0026d=identicon",
		      "web_url": "https://gitlab.com/michaellihs"
		    },
		    "name": "test",
		    "name_with_namespace": "Michael Lihs / test",
		    "path": "test",
		    "path_with_namespace": "michaellihs/test",
		    "container_registry_enabled": true,
		    "issues_enabled": true,
		    "merge_requests_enabled": true,
		    "wiki_enabled": true,
		    "builds_enabled": true,
		    "snippets_enabled": false,
		    "created_at": "2017-03-08T22:55:11.313Z",
		    "last_activity_at": "2017-03-08T22:55:11.313Z",
		    "shared_runners_enabled": true,
		    "lfs_enabled": true,
		    "creator_id": 1150266,
		    "namespace": {
		      "id": 1368278,
		      "name": "michaellihs",
		      "path": "michaellihs",
		      "kind": "user"
		    },
		    "avatar_url": null,
		    "star_count": 0,
		    "forks_count": 0,
		    "open_issues_count": 0,
		    "public_builds": true,
		    "shared_with_groups": [],
		    "only_allow_merge_if_build_succeeds": false,
		    "request_access_enabled": false,
		    "only_allow_merge_if_all_discussions_are_resolved": false,
		    "approvals_before_merge": 0,
		    "permissions": {
		      "project_access": {
			"access_level": 40,
			"notification_level": 3
		      },
		      "group_access": null
		    }
		  }`), project)
		Expect(project.Name).To(Equal("test"))
		Expect(project.ID).To(Equal(2859355))
		Expect(project.Namespace.ID).To(Equal(1368278))
		Expect(project.Namespace.Name).To(Equal("michaellihs"))
		Expect(project.Namespace.Path).To(Equal("michaellihs"))
		Expect(project.Namespace.Kind).To(Equal("user"))
	})

	It("can be unmarshaled from JSON with group_access", func() {
		project := new(Project)
		json.Unmarshal([]byte(`{
		    "id": 2859355,
		    "description": null,
		    "default_branch": null,
		    "tag_list": [],
		    "public": false,
		    "archived": false,
		    "visibility_level": 0,
		    "ssh_url_to_repo": "git@gitlab.com:michaellihs/test.git",
		    "http_url_to_repo": "https://gitlab.com/michaellihs/test.git",
		    "web_url": "https://gitlab.com/michaellihs/test",
		    "owner": {
		      "name": "Michael Lihs",
		      "username": "michaellihs",
		      "id": 1150266,
		      "state": "active",
		      "avatar_url": "https://secure.gravatar.com/avatar/4e1d79b0859761cf6de8f560784226cc?s=80\u0026d=identicon",
		      "web_url": "https://gitlab.com/michaellihs"
		    },
		    "name": "test",
		    "name_with_namespace": "Michael Lihs / test",
		    "path": "test",
		    "path_with_namespace": "michaellihs/test",
		    "container_registry_enabled": true,
		    "issues_enabled": true,
		    "merge_requests_enabled": true,
		    "wiki_enabled": true,
		    "builds_enabled": true,
		    "snippets_enabled": false,
		    "created_at": "2017-03-08T22:55:11.313Z",
		    "last_activity_at": "2017-03-08T22:55:11.313Z",
		    "shared_runners_enabled": true,
		    "lfs_enabled": true,
		    "creator_id": 1150266,
		    "namespace": {
		      "id": 1368278,
		      "name": "michaellihs",
		      "path": "michaellihs",
		      "kind": "user"
		    },
		    "avatar_url": null,
		    "star_count": 0,
		    "forks_count": 0,
		    "open_issues_count": 0,
		    "public_builds": true,
		    "shared_with_groups": [],
		    "only_allow_merge_if_build_succeeds": false,
		    "request_access_enabled": false,
		    "only_allow_merge_if_all_discussions_are_resolved": false,
		    "approvals_before_merge": 0,
		    "permissions": {
		      "group_access": {
			"access_level": 40,
			"notification_level": 3
		      },
		      "project_access": null
		    }
		  }`), project)
		Expect(project.Name).To(Equal("test"))
		Expect(project.ID).To(Equal(2859355))
		Expect(project.Namespace.ID).To(Equal(1368278))
		Expect(project.Namespace.Name).To(Equal("michaellihs"))
		Expect(project.Namespace.Path).To(Equal("michaellihs"))
		Expect(project.Namespace.Kind).To(Equal("user"))
	})

	It("can be unmarshaled from JSON with group_access and project_access", func() {
		project := new(Project)
		json.Unmarshal([]byte(`{
		    "id": 2859355,
		    "description": null,
		    "default_branch": null,
		    "tag_list": [],
		    "public": false,
		    "archived": false,
		    "visibility_level": 0,
		    "ssh_url_to_repo": "git@gitlab.com:michaellihs/test.git",
		    "http_url_to_repo": "https://gitlab.com/michaellihs/test.git",
		    "web_url": "https://gitlab.com/michaellihs/test",
		    "owner": {
		      "name": "Michael Lihs",
		      "username": "michaellihs",
		      "id": 1150266,
		      "state": "active",
		      "avatar_url": "https://secure.gravatar.com/avatar/4e1d79b0859761cf6de8f560784226cc?s=80\u0026d=identicon",
		      "web_url": "https://gitlab.com/michaellihs"
		    },
		    "name": "test",
		    "name_with_namespace": "Michael Lihs / test",
		    "path": "test",
		    "path_with_namespace": "michaellihs/test",
		    "container_registry_enabled": true,
		    "issues_enabled": true,
		    "merge_requests_enabled": true,
		    "wiki_enabled": true,
		    "builds_enabled": true,
		    "snippets_enabled": false,
		    "created_at": "2017-03-08T22:55:11.313Z",
		    "last_activity_at": "2017-03-08T22:55:11.313Z",
		    "shared_runners_enabled": true,
		    "lfs_enabled": true,
		    "creator_id": 1150266,
		    "namespace": {
		      "id": 1368278,
		      "name": "michaellihs",
		      "path": "michaellihs",
		      "kind": "user"
		    },
		    "avatar_url": null,
		    "star_count": 0,
		    "forks_count": 0,
		    "open_issues_count": 0,
		    "public_builds": true,
		    "shared_with_groups": [],
		    "only_allow_merge_if_build_succeeds": false,
		    "request_access_enabled": false,
		    "only_allow_merge_if_all_discussions_are_resolved": false,
		    "approvals_before_merge": 0,
		    "permissions": {
		      "group_access": {
			"access_level": 40,
			"notification_level": 3
		      },
		      "project_access": null
		    }
		  }`), project)
		Expect(project.Name).To(Equal("test"))
		Expect(project.ID).To(Equal(2859355))
		Expect(project.Namespace.ID).To(Equal(1368278))
		Expect(project.Namespace.Name).To(Equal("michaellihs"))
		Expect(project.Namespace.Path).To(Equal("michaellihs"))
		Expect(project.Namespace.Kind).To(Equal("user"))
	})

})
