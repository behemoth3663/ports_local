PORTNAME=	clash
DISTVERSIONPREFIX=	v
DISTVERSION=	1.16.0
CATEGORIES=	local net

MAINTAINER=	yuri@FreeBSD.org
COMMENT=	Rule-based tunnel in Go
WWW=		https://github.com/${GH_ACCOUNT}/${PORTNAME}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	Dreamacro

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/constant.Version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/constant.BuildTime=${%FT%TZ:L:gmtime}'"

pre-patch: .PHONY .SILENT
	${CP} ${WRKSRC}/listener/redir/tcp_linux.go ${WRKSRC}/listener/redir/tcp_freebsd.go

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
