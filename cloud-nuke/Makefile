#

PORTNAME=	cloud-nuke
PORTVERSION=	0.16.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool to delete all resources in an AWS account

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	gruntwork-io

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract=	501:pre-post-extract

pre-post-extract: .PHONY
	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
