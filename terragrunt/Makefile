PORTNAME=	terragrunt
DISTVERSIONPREFIX=	v
DISTVERSION=	0.64.0
#PORTREVISION=	1
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

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-googleapis:google-cloud-go:v0.2.3:googleapis_google_cloud_go_1/vendor/cloud.google.com/go/compute/metadata
#+googleapis:google-cloud-go:compute/metadata/v0.2.3:googleapis_google_cloud_go_1/vendor/cloud.google.com/go/compute/metadata
GH_TUPLE:=	${GH_TUPLE:C,^(googleapis:google-cloud-go):(.*/)(compute/metadata|auth/oauth2adapt)(.*),\1:\3/\2\3\4,1}

#-hashicorp:vault:v1.10.0:hashicorp_vault_api
#+hashicorp:vault:api/v1.10.0:hashicorp_vault_api
GH_TUPLE:=	${GH_TUPLE:S|^hashicorp:vault:|hashicorp:vault:api/|1}

.include <bsd.port.mk>
