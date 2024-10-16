PORTNAME=	terragrunt
DISTVERSIONPREFIX=	v
DISTVERSION=	0.68.3
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	dutra@FreeBSD.org
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

M2T_OFFLINE=	true
.export	M2T_OFFLINE

_OPTIONS_extract=	600:pre-post-extract
pre-post-extract: .PHONY
	${RM} ${WRKSRC}/terraform/getproviders/lock_test.go
	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	test -f ${WRKSRC}/vendor/modules.txt || test ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
