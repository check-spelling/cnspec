// Code generated by protoc-gen-rangerrpc version DO NOT EDIT.
// source: cnspec_policy.proto

package policy

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	ranger "go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/metadata"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

// service interface definition

type PolicyHub interface {
	SetBundle(context.Context, *Bundle) (*Empty, error)
	ValidateBundle(context.Context, *Bundle) (*Empty, error)
	GetBundle(context.Context, *Mrn) (*Bundle, error)
	GetPolicy(context.Context, *Mrn) (*Policy, error)
	DeletePolicy(context.Context, *Mrn) (*Empty, error)
	GetPolicyFilters(context.Context, *Mrn) (*Mqueries, error)
	List(context.Context, *ListReq) (*Policies, error)
	DefaultPolicies(context.Context, *DefaultPoliciesReq) (*URLs, error)
}

// client implementation

type PolicyHubClient struct {
	ranger.Client
	httpclient ranger.HTTPClient
	prefix     string
}

func NewPolicyHubClient(addr string, client ranger.HTTPClient, plugins ...ranger.ClientPlugin) (*PolicyHubClient, error) {
	base, err := url.Parse(ranger.SanitizeUrl(addr))
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("./PolicyHub")
	if err != nil {
		return nil, err
	}

	serviceClient := &PolicyHubClient{
		httpclient: client,
		prefix:     base.ResolveReference(u).String(),
	}
	serviceClient.AddPlugins(plugins...)
	return serviceClient, nil
}
func (c *PolicyHubClient) SetBundle(ctx context.Context, in *Bundle) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/SetBundle"}, ""), in, out)
	return out, err
}
func (c *PolicyHubClient) ValidateBundle(ctx context.Context, in *Bundle) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/ValidateBundle"}, ""), in, out)
	return out, err
}
func (c *PolicyHubClient) GetBundle(ctx context.Context, in *Mrn) (*Bundle, error) {
	out := new(Bundle)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetBundle"}, ""), in, out)
	return out, err
}
func (c *PolicyHubClient) GetPolicy(ctx context.Context, in *Mrn) (*Policy, error) {
	out := new(Policy)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetPolicy"}, ""), in, out)
	return out, err
}
func (c *PolicyHubClient) DeletePolicy(ctx context.Context, in *Mrn) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/DeletePolicy"}, ""), in, out)
	return out, err
}
func (c *PolicyHubClient) GetPolicyFilters(ctx context.Context, in *Mrn) (*Mqueries, error) {
	out := new(Mqueries)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetPolicyFilters"}, ""), in, out)
	return out, err
}
func (c *PolicyHubClient) List(ctx context.Context, in *ListReq) (*Policies, error) {
	out := new(Policies)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/List"}, ""), in, out)
	return out, err
}
func (c *PolicyHubClient) DefaultPolicies(ctx context.Context, in *DefaultPoliciesReq) (*URLs, error) {
	out := new(URLs)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/DefaultPolicies"}, ""), in, out)
	return out, err
}

// server implementation

type PolicyHubServerOption func(s *PolicyHubServer)

func WithUnknownFieldsForPolicyHubServer() PolicyHubServerOption {
	return func(s *PolicyHubServer) {
		s.allowUnknownFields = true
	}
}

func NewPolicyHubServer(handler PolicyHub, opts ...PolicyHubServerOption) http.Handler {
	srv := &PolicyHubServer{
		handler: handler,
	}

	for i := range opts {
		opts[i](srv)
	}

	service := ranger.Service{
		Name: "PolicyHub",
		Methods: map[string]ranger.Method{
			"SetBundle":        srv.SetBundle,
			"ValidateBundle":   srv.ValidateBundle,
			"GetBundle":        srv.GetBundle,
			"GetPolicy":        srv.GetPolicy,
			"DeletePolicy":     srv.DeletePolicy,
			"GetPolicyFilters": srv.GetPolicyFilters,
			"List":             srv.List,
			"DefaultPolicies":  srv.DefaultPolicies,
		},
	}
	return ranger.NewRPCServer(&service)
}

type PolicyHubServer struct {
	handler            PolicyHub
	allowUnknownFields bool
}

func (p *PolicyHubServer) SetBundle(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Bundle
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.SetBundle(ctx, &req)
}
func (p *PolicyHubServer) ValidateBundle(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Bundle
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.ValidateBundle(ctx, &req)
}
func (p *PolicyHubServer) GetBundle(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Mrn
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetBundle(ctx, &req)
}
func (p *PolicyHubServer) GetPolicy(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Mrn
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetPolicy(ctx, &req)
}
func (p *PolicyHubServer) DeletePolicy(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Mrn
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.DeletePolicy(ctx, &req)
}
func (p *PolicyHubServer) GetPolicyFilters(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Mrn
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetPolicyFilters(ctx, &req)
}
func (p *PolicyHubServer) List(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req ListReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.List(ctx, &req)
}
func (p *PolicyHubServer) DefaultPolicies(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req DefaultPoliciesReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.DefaultPolicies(ctx, &req)
}

// service interface definition

type PolicyResolver interface {
	Assign(context.Context, *PolicyAssignment) (*Empty, error)
	Unassign(context.Context, *PolicyAssignment) (*Empty, error)
	Resolve(context.Context, *ResolveReq) (*ResolvedPolicy, error)
	UpdateAssetJobs(context.Context, *UpdateAssetJobsReq) (*Empty, error)
	ResolveAndUpdateJobs(context.Context, *UpdateAssetJobsReq) (*ResolvedPolicy, error)
	GetResolvedPolicy(context.Context, *Mrn) (*ResolvedPolicy, error)
	StoreResults(context.Context, *StoreResultsReq) (*Empty, error)
	GetReport(context.Context, *EntityScoreReq) (*Report, error)
	GetScore(context.Context, *EntityScoreReq) (*Report, error)
	SynchronizeAssets(context.Context, *SynchronizeAssetsReq) (*SynchronizeAssetsResp, error)
	PurgeAssets(context.Context, *PurgeAssetsRequest) (*PurgeAssetsConfirmation, error)
}

// client implementation

type PolicyResolverClient struct {
	ranger.Client
	httpclient ranger.HTTPClient
	prefix     string
}

func NewPolicyResolverClient(addr string, client ranger.HTTPClient, plugins ...ranger.ClientPlugin) (*PolicyResolverClient, error) {
	base, err := url.Parse(ranger.SanitizeUrl(addr))
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("./PolicyResolver")
	if err != nil {
		return nil, err
	}

	serviceClient := &PolicyResolverClient{
		httpclient: client,
		prefix:     base.ResolveReference(u).String(),
	}
	serviceClient.AddPlugins(plugins...)
	return serviceClient, nil
}
func (c *PolicyResolverClient) Assign(ctx context.Context, in *PolicyAssignment) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Assign"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) Unassign(ctx context.Context, in *PolicyAssignment) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Unassign"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) Resolve(ctx context.Context, in *ResolveReq) (*ResolvedPolicy, error) {
	out := new(ResolvedPolicy)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Resolve"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) UpdateAssetJobs(ctx context.Context, in *UpdateAssetJobsReq) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/UpdateAssetJobs"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) ResolveAndUpdateJobs(ctx context.Context, in *UpdateAssetJobsReq) (*ResolvedPolicy, error) {
	out := new(ResolvedPolicy)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/ResolveAndUpdateJobs"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) GetResolvedPolicy(ctx context.Context, in *Mrn) (*ResolvedPolicy, error) {
	out := new(ResolvedPolicy)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetResolvedPolicy"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) StoreResults(ctx context.Context, in *StoreResultsReq) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/StoreResults"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) GetReport(ctx context.Context, in *EntityScoreReq) (*Report, error) {
	out := new(Report)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetReport"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) GetScore(ctx context.Context, in *EntityScoreReq) (*Report, error) {
	out := new(Report)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetScore"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) SynchronizeAssets(ctx context.Context, in *SynchronizeAssetsReq) (*SynchronizeAssetsResp, error) {
	out := new(SynchronizeAssetsResp)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/SynchronizeAssets"}, ""), in, out)
	return out, err
}
func (c *PolicyResolverClient) PurgeAssets(ctx context.Context, in *PurgeAssetsRequest) (*PurgeAssetsConfirmation, error) {
	out := new(PurgeAssetsConfirmation)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/PurgeAssets"}, ""), in, out)
	return out, err
}

// server implementation

type PolicyResolverServerOption func(s *PolicyResolverServer)

func WithUnknownFieldsForPolicyResolverServer() PolicyResolverServerOption {
	return func(s *PolicyResolverServer) {
		s.allowUnknownFields = true
	}
}

func NewPolicyResolverServer(handler PolicyResolver, opts ...PolicyResolverServerOption) http.Handler {
	srv := &PolicyResolverServer{
		handler: handler,
	}

	for i := range opts {
		opts[i](srv)
	}

	service := ranger.Service{
		Name: "PolicyResolver",
		Methods: map[string]ranger.Method{
			"Assign":               srv.Assign,
			"Unassign":             srv.Unassign,
			"Resolve":              srv.Resolve,
			"UpdateAssetJobs":      srv.UpdateAssetJobs,
			"ResolveAndUpdateJobs": srv.ResolveAndUpdateJobs,
			"GetResolvedPolicy":    srv.GetResolvedPolicy,
			"StoreResults":         srv.StoreResults,
			"GetReport":            srv.GetReport,
			"GetScore":             srv.GetScore,
			"SynchronizeAssets":    srv.SynchronizeAssets,
			"PurgeAssets":          srv.PurgeAssets,
		},
	}
	return ranger.NewRPCServer(&service)
}

type PolicyResolverServer struct {
	handler            PolicyResolver
	allowUnknownFields bool
}

func (p *PolicyResolverServer) Assign(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req PolicyAssignment
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Assign(ctx, &req)
}
func (p *PolicyResolverServer) Unassign(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req PolicyAssignment
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Unassign(ctx, &req)
}
func (p *PolicyResolverServer) Resolve(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req ResolveReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Resolve(ctx, &req)
}
func (p *PolicyResolverServer) UpdateAssetJobs(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req UpdateAssetJobsReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.UpdateAssetJobs(ctx, &req)
}
func (p *PolicyResolverServer) ResolveAndUpdateJobs(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req UpdateAssetJobsReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.ResolveAndUpdateJobs(ctx, &req)
}
func (p *PolicyResolverServer) GetResolvedPolicy(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Mrn
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetResolvedPolicy(ctx, &req)
}
func (p *PolicyResolverServer) StoreResults(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req StoreResultsReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.StoreResults(ctx, &req)
}
func (p *PolicyResolverServer) GetReport(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req EntityScoreReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetReport(ctx, &req)
}
func (p *PolicyResolverServer) GetScore(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req EntityScoreReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetScore(ctx, &req)
}
func (p *PolicyResolverServer) SynchronizeAssets(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req SynchronizeAssetsReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.SynchronizeAssets(ctx, &req)
}
func (p *PolicyResolverServer) PurgeAssets(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req PurgeAssetsRequest
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.PurgeAssets(ctx, &req)
}
