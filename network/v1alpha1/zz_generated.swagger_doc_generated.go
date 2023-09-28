package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_DNSNameResolver = map[string]string{
	"":         "DNSNameResolver stores the DNS name resolution information of a DNS name. It can be enabled by the TechPreviewNoUpgrade feature set. It can also be enabled by the feature gate DNSNameResolver when using CustomNoUpgrade feature set.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec is the specification of the desired behavior of the DNSNameResolver.",
	"status":   "status is the most recently observed status of the DNSNameResolver.",
}

func (DNSNameResolver) SwaggerDoc() map[string]string {
	return map_DNSNameResolver
}

var map_DNSNameResolverList = map[string]string{
	"":         "DNSNameResolverList contains a list of DNSNameResolvers.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (DNSNameResolverList) SwaggerDoc() map[string]string {
	return map_DNSNameResolverList
}

var map_DNSNameResolverResolvedAddress = map[string]string{
	"ip":             "ip is an IP address associated with the dnsName. The validity of the IP address expires after lastLookupTime + ttlSeconds. To refresh the information, a DNS lookup will be performed upon the expiration of the IP address's validity. If the information is not refreshed then it will be removed with a grace period after the expiration of the IP address's validity.",
	"ttlSeconds":     "ttlSeconds is the time-to-live value of the IP address. The validity of the IP address expires after lastLookupTime + ttlSeconds. On a successful DNS lookup the value of this field will be updated with the current time-to-live value. If the information is not refreshed then it will be removed with a grace period after the expiration of the IP address's validity.",
	"lastLookupTime": "lastLookupTime is the timestamp when the last DNS lookup was completed successfully. The validity of the IP address expires after lastLookupTime + ttlSeconds. The value of this field will be updated to the current time on a successful DNS lookup. If the information is not refreshed then it will be removed with a grace period after the expiration of the IP address's validity.",
}

func (DNSNameResolverResolvedAddress) SwaggerDoc() map[string]string {
	return map_DNSNameResolverResolvedAddress
}

var map_DNSNameResolverResolvedName = map[string]string{
	"":                   "DNSNameResolverResolvedName describes the details of a resolved DNS name.",
	"conditions":         "conditions provide information about the state of the DNS name. Known .status.conditions.type is: \"Degraded\". \"Degraded\" is true when the last resolution failed for the DNS name, and false otherwise.",
	"dnsName":            "dnsName is the resolved DNS name matching the name field of DNSNameResolverSpec. This field can store both regular and wildcard DNS names which match the spec.name field. When the spec.name field contains a regular DNS name, this field will store the same regular DNS name after it is successfully resolved. When the spec.name field contains a wildcard DNS name, each resolvedName.dnsName will store the regular DNS names which match the wildcard DNS name and have been successfully resolved. If the wildcard DNS name can also be successfully resolved, then this field will store the wildcard DNS name as well.",
	"resolvedAddresses":  "resolvedAddresses gives the list of associated IP addresses and their corresponding TTLs and last lookup times for the dnsName.",
	"resolutionFailures": "resolutionFailures keeps the count of how many consecutive times the DNS resolution failed for the dnsName. If the DNS resolution succeeds then the field will be set to zero. Upon every failure, the value of the field will be incremented by one. The details about the DNS name will be removed, if the value of resolutionFailures reaches 5 and the TTL of all the associated IP addresses have expired.",
}

func (DNSNameResolverResolvedName) SwaggerDoc() map[string]string {
	return map_DNSNameResolverResolvedName
}

var map_DNSNameResolverSpec = map[string]string{
	"":     "DNSNameResolverSpec is a desired state description of DNSNameResolver.",
	"name": "name is the DNS name for which the DNS name resolution information will be stored. For a regular DNS name, only the DNS name resolution information of the regular DNS name will be stored. For a wildcard DNS name, the DNS name resolution information of all the DNS names that match the wildcard DNS name will be stored. For a wildcard DNS name, the '*' will match only one label. Additionally, only a single '*' can be used at the beginning of the wildcard DNS name. For example, '*.example.com.' will match 'sub1.example.com.' but won't match 'sub2.sub1.example.com.'",
}

func (DNSNameResolverSpec) SwaggerDoc() map[string]string {
	return map_DNSNameResolverSpec
}

var map_DNSNameResolverStatus = map[string]string{
	"":              "DNSNameResolverStatus defines the observed status of DNSNameResolver.",
	"resolvedNames": "resolvedNames contains a list of matching DNS names and their corresponding IP addresses along with their TTL and last DNS lookup times.",
}

func (DNSNameResolverStatus) SwaggerDoc() map[string]string {
	return map_DNSNameResolverStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
