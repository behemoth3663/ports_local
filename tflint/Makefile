# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	tflint
PORTVERSION=	0.55.1
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	dutra@FreeBSD.org
COMMENT=	Terraform Linter
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-linters

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

PLIST_FILES=	bin/${PORTNAME}

#It seems to be hard to make this work
#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

pre-deps post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-GoogleCloudPlatform:opentelemetry-operations-go:d637680bde0c:googlecloudplatform_opentelemetry_operations_go_gcp
#-census-instrumentation:opencensus-go:v0.24.0:census_instrumentation_opencensus_go/vendor/go.opencensus.io
#-open-telemetry:opentelemetry-go-contrib:v0.54.0:open_telemetry_opentelemetry_go_contrib_1/vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc
#-opentracing:opentracing-go:v1.2.0:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go
GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:*}
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:*}
GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:*}
GH_TUPLE:=	${GH_TUPLE:Nopentracing\:*}

#-google:go-genproto:69823020774d:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:69823020774d:google_go_genproto_2
# google:go-genproto:e639e219e697:google_go_genproto/vendor/google.golang.org/genproto
#-google:go-genproto:e6fa225c2576:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:e6fa225c2576:google_go_genproto_1
GH_TUPLE:=	${GH_TUPLE:C|_[[:digit:]]/vendor/google.golang.org/genproto/googleapis/|_|}

#make -V GH_TUPLE | tr ' ' '\n' | i googleapis:google-cloud-go
#	googleapis:google-cloud-go:v0.112.1:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-	googleapis:google-cloud-go:v0.13.0:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+	googleapis:google-cloud-go:auth/v0.13.0:googleapis_google_cloud_go_auth
#-	googleapis:google-cloud-go:v0.6.0:googleapis_google_cloud_go_2/vendor/cloud.google.com/go/compute/metadata
#+	googleapis:google-cloud-go:compute/metadata/v0.6.0:googleapis_google_cloud_go_compute_metadata
GH_TUPLE:=	${GH_TUPLE:N*\:googleapis_google_cloud_go_[b-z]*}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(auth)$|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/0-9]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2\3_\4|}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/cloud.google.com/go/*}

GH_TUPLE+=	hashicorp:go-version:v1.7.0:hashicorp_go_version/vendor/github.com/hashicorp/go-version
GH_TUPLE+=	spf13:cast:v1.7.0:spf13_cast/vendor/github.com/spf13/cast

.include <bsd.port.mk>
