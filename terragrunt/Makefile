# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	terragrunt
DISTVERSIONPREFIX=	v
DISTVERSION=	0.35.8
CATEGORIES=	sysutils

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

#${LN} -s ${WRKDIR}/go-commons-0.7.0 ${WRKDIR}/gruntwork-cli-0.7.0
pre-post-extract: .PHONY
	src_dst="$$(${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | ${SED} -r 's|(.*/)(go-commons-)(.*)|\1\2\3 \1gruntwork-cli-\3|')"; \
	test -z "$${src_dst}" || ${LN} -s $${src_dst}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
