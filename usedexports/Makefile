PORTNAME=	usedexports
PORTVERSION=	0.0.1
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Find in Go exported variables that could be unexported
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

#USES=		go:modules
USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	jgautheron
GH_TAGNAME=	69abd62

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

do-build: .PHONY
	(cd ${GO_WRKSRC} && \
		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} build ${GO_BUILDFLAGS} \
			-o ${GO_WRKDIR_BIN}/${PORTNAME})

.include <bsd.port.mk>
