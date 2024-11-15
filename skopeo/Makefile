PORTNAME=	skopeo
PORTVERSION=	1.17.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Utility for various operations on container images and repositories
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules,1.23
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	containers

PLIST_FILES=	bin/${PORTNAME}

#-extldflags '-static' causes ld: error: unable to find library -lgpgme
GO_BUILDFLAGS+=	-ldflags="-s -w"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-moby:moby:v0.8.2:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.8.2:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/)(docker-credential-helpers)$|docker:\2\1\2|1}

#-open-telemetry:opentelemetry-go-contrib:v0.45.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
#+open-telemetry:opentelemetry-go-contrib:instrumentation/net/http/otelhttp/v0.45.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
GH_TUPLE:=	${GH_TUPLE:C|^(open-telemetry:opentelemetry-go-contrib:)(.*/)(instrumentation/net/http/otelhttp)$|\1\3/\2\3|1}

#-open-telemetry:opentelemetry-proto-go:v1.0.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto/otlp
#+open-telemetry:opentelemetry-proto-go:otlp/v1.0.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto/otlp
#GH_TUPLE:=	${GH_TUPLE:C|^(open-telemetry:opentelemetry-proto-go:)(.*/)(otlp)$|\1\3/\2\3|1}

.include <bsd.port.mk>
