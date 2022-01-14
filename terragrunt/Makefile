# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	terragrunt
DISTVERSIONPREFIX=	v
DISTVERSION=	0.35.19
CATEGORIES=	local sysutils

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Terragrunt is a thin wrapper for Terraform

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	gruntwork-io

GO_BUILDFLAGS=	-ldflags="-s -w -X main.VERSION=${DISTVERSIONPREFIX}${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract=	501:pre-post-extract

pre-post-extract: .PHONY
	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
