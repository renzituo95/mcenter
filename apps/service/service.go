package service

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/infraboard/mcenter/apps/service/provider/gitlab"
)

func (s *ServiceSet) Len() int {
	return len(s.Items)
}

func (s *ServiceSet) UpdateFromGitProject(p *gitlab.Project, tk string) {
	svc := s.GetServiceByGitSshUrl(p.GitSshUrl)
	if svc == nil {
		svc = NewServiceFromProject(p)
		svc.Spec.Repository.Token = tk
		s.Add(svc)
	}
}

func (s *ServiceSet) UpdateScope(domain, namespace string) {
	for i := range s.Items {
		item := s.Items[i]
		item.Spec.Domain = domain
		item.Spec.Namespace = namespace
	}
}

func (s *ServiceSet) GetServiceByGitSshUrl(gitSshUrl string) *Service {
	for i := range s.Items {
		item := s.Items[i]
		if item.GetRepositorySshUrl() == gitSshUrl {
			return item
		}
	}

	return nil
}

func NewServiceFromProject(p *gitlab.Project) *Service {
	svc := NewDefaultService()
	spec := svc.Spec
	spec.Name = p.Name
	spec.Logo = p.AvatarURL
	spec.Description = p.Description
	spec.Repository.ProjectId = p.IdToString()
	spec.Repository.SshUrl = p.GitSshUrl
	spec.Repository.HttpUrl = p.GitHttpUrl
	spec.Repository.Namespace = p.NamespacePath
	spec.Repository.WebUrl = p.WebURL
	spec.Repository.CreatedAt = p.CreatedAt.Unix()
	spec.Repository.EnableHook = true
	spec.Repository.HookConfig = gitlab.NewGitLabWebHook(
		"自动填充服务的Id",
	).ToJson()
	return svc
}

func (s *Service) GetRepositorySshUrl() string {
	if s.Spec.Repository != nil {
		return s.Spec.Repository.SshUrl
	}

	return ""
}

func NewRepository() *Repository {
	return &Repository{
		EnableHook: true,
	}
}

func (r *Repository) SetLanguage(v LANGUAGE) {
	r.Language = &v
}

func (r *Repository) ProjectIdToInt64() int64 {
	pid, _ := strconv.ParseInt(r.ProjectId, 10, 64)
	return pid
}

func (r *Repository) HookIdToInt64() int64 {
	pid, _ := strconv.ParseInt(r.HookId, 10, 64)
	return pid
}

func (r *Repository) MakeGitlabConfig() (*gitlab.Config, error) {
	conf := gitlab.NewDefaultConfig()
	addr, err := r.HostAddress()
	if err != nil {
		return nil, err
	}
	conf.Address = addr
	conf.PrivateToken = r.Token
	return conf, nil
}

func (r *Repository) HostAddress() (string, error) {
	u, err := url.Parse(r.WebUrl)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host), nil
}