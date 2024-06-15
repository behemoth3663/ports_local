PORTNAME=	terragrunt
DISTVERSIONPREFIX=	v
DISTVERSION=	0.59.3
CATEGORIES=	local sysutils

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Terragrunt is a thin wrapper for Terraform
WWW=		https://terragrunt.gruntwork.io/

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	gruntwork-io

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/go-commons/version.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract=	600:pre-post-extract
pre-post-extract: .PHONY
	${RM} ${WRKSRC}/terraform/getproviders/lock_test.go
	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

#-		open-telemetry:opentelemetry-go:v1.23.1:open_telemetry_opentelemetry_go/vendor/go.opentelemetry.io/otel
#-		open-telemetry:opentelemetry-proto-go:v1.1.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto/otlp
#GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:opentelemetry-*}
#post-patch: .PHONY .SILENT
#	if cd ${PATCHDIR} 2>/dev/null; then \
#		${FIND} . -mindepth 1 -maxdepth 1 -type d -print0 | \
#		${TAR} --null --files-from - -cf- | \
#		${TAR} -C ${WRKSRC} -xf-; \
#	fi

pre-deps: .PHONY
	if [ -d '${PATCHDIR}_' ]; then \
		${RM} -r -- ${PATCHDIR}; \
	else \
		${MV} ${PATCHDIR} ${PATCHDIR}_; \
	fi

pre-deps-all: .PHONY

post-deps-all: .PHONY
	test ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-		googleapis:google-cloud-go:v0.2.3:googleapis_google_cloud_go_1/vendor/cloud.google.com/go/compute/metadata
#+		googleapis:google-cloud-go:compute/metadata/v0.2.3:googleapis_google_cloud_go_1/vendor/cloud.google.com/go/compute/metadata
GH_TUPLE:=	${GH_TUPLE:C|^(googleapis:google-cloud-go):(.*/compute/metadata.*)|\1:compute/metadata/\2|1}

#-       hashicorp:vault:v1.10.0:hashicorp_vault_api
#+       hashicorp:vault:api/v1.10.0:hashicorp_vault_api
GH_TUPLE:=	${GH_TUPLE:S|^hashicorp:vault:|hashicorp:vault:api/|1}

.include <bsd.port.mk>
