#

PORTNAME=	dyff
DISTVERSIONPREFIX=	v
DISTVERSION=	1.4.7
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	diff tool for YAML files, and sometimes JSON

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	homeport

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS=	-ldflags="-s -w \
		-X 'github.com/${GH_ACCOUNT}/${PORTNAME}/internal/cmd.version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

#_OPTIONS_extract=	501:pre-post-extract

#pre-post-extract: .PHONY
#	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
#		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
#	done

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
