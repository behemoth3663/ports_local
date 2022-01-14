# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	gox
DISTVERSIONPREFIX=	v
DISTVERSION=	1.0.1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Go cross compile tool

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mitchellh

#GO_BUILDFLAGS=	-ldflags="-s -w -X main.VERSION=${DISTVERSIONPREFIX}${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

#_OPTIONS_extract=	501:pre-post-extract
#pre-post-extract: .PHONY
#	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
#		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
#	done

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
