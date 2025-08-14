PORTNAME=	skopeo
PORTVERSION=	1.20.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Utility for various operations on container images and repositories
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	containers

PLIST_FILES=	bin/${PORTNAME}

#-extldflags '-static' causes ld: error: unable to find library -lgpgme
GO_BUILDFLAGS+=	-ldflags="-s -w"

#1.19.0
#===>  Building skopeo from ./cmd/skopeo
#vendor/github.com/go-openapi/runtime/client/opentracing.go:8:2: cannot find module providing package github.com/opentracing/opentracing-go: import lookup disabled by -mod=vendor
#vendor/github.com/go-openapi/runtime/client/opentracing.go:9:2: cannot find module providing package github.com/opentracing/opentracing-go/ext: import lookup disabled by -mod=vendor
#vendor/github.com/go-openapi/runtime/client/opentracing.go:10:2: cannot find module providing package github.com/opentracing/opentracing-go/log: import lookup disabled by -mod=vendor
#vendor/github.com/go-openapi/runtime/client/opentelemetry.go:10:2: cannot find module providing package go.opentelemetry.io/otel: import lookup disabled by -mod=vendor
#vendor/github.com/go-openapi/runtime/client/opentelemetry.go:11:2: cannot find module providing package go.opentelemetry.io/otel/attribute: import lookup disabled by -mod=vendor
#vendor/github.com/go-openapi/runtime/client/opentelemetry.go:12:2: cannot find module providing package go.opentelemetry.io/otel/codes: import lookup disabled by -mod=vendor
#vendor/github.com/go-openapi/runtime/client/opentelemetry.go:13:2: cannot find module providing package go.opentelemetry.io/otel/propagation: import lookup disabled by -mod=vendor
#vendor/github.com/go-openapi/runtime/client/opentelemetry.go:14:2: cannot find module providing package go.opentelemetry.io/otel/semconv/v1.17.0: import lookup disabled by -mod=vendor
#vendor/github.com/go-openapi/runtime/client/opentelemetry.go:15:2: cannot find module providing package go.opentelemetry.io/otel/semconv/v1.17.0/httpconv: import lookup disabled by -mod=vendor
#vendor/github.com/go-openapi/runtime/client/opentelemetry.go:16:2: cannot find module providing package go.opentelemetry.io/otel/trace: import lookup disabled by -mod=vendor
#vendor/github.com/docker/docker/client/client.go:61:2: cannot find module providing package go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: import lookup disabled by -mod=vendor
#*** Error code 1
#_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	@${RM} -r ${WRKSRC}/vendor/go.opencensus.io
	@${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	@${RM} -r ${WRKSRC}/vendor/github.com/opentracing

pre-deps post-deps-all pre-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-moby:moby:v0.8.2:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.8.2:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/)(docker-credential-helpers)$|docker:\2\1\2|1}

#-open-telemetry:opentelemetry-go-contrib:v0.45.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
#-open-telemetry:opentelemetry-go-instrumentation:v1.1.0:open_telemetry_opentelemetry_go_instrumentation/vendor/go.opentelemetry.io/auto/sdk
#-open-telemetry:opentelemetry-proto-go:v1.0.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto/otlp
GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:opentelemetry-go*}

#-opentracing:opentracing-go:v1.2.0:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go
GH_TUPLE:=	${GH_TUPLE:Nopentracing\:opentracing-go\:*}

#-census-instrumentation:opencensus-go:v0.24.0:census_instrumentation_opencensus_go/vendor/go.opencensus.io
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:opencensus-go\:*}

.include <bsd.port.mk>
