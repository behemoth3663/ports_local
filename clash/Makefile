PORTNAME=	clash
DISTVERSIONPREFIX=	v
DISTVERSION=	1.8.0
CATEGORIES=	net

MAINTAINER=	yuri@FreeBSD.org
COMMENT=	Rule-based tunnel in Go

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

#GO_MODULE=	github.com/Dreamacro/clash
USE_GITHUB=	yes
GH_ACCOUNT=	Dreamacro

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/constant.Version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/constant.BuildTime=${%FT%TZ:L:gmtime}'"

pre-patch: .PHONY .SILENT
	${CP} ${WRKSRC}/listener/redir/tcp_linux.go ${WRKSRC}/listener/redir/tcp_freebsd.go

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
