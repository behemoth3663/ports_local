PORTNAME=	clash
DISTVERSIONPREFIX=	v
DISTVERSION=	1.7.1
CATEGORIES=	net

MAINTAINER=	yuri@FreeBSD.org
COMMENT=	Rule-based tunnel in Go

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

GO_MODULE=	github.com/Dreamacro/clash

PLIST_FILES=	bin/${PORTNAME}

PORTSCOUT=	limit:^[0-9\.]*$$

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X '${GO_MODULE}/constant.Version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X '${GO_MODULE}/constant.BuildTime=${%FT%TZ:L:gmtime}'"

pre-patch: .PHONY .SILENT
	${CP} ${WRKSRC}/listener/redir/tcp_linux.go ${WRKSRC}/listener/redir/tcp_freebsd.go

.include <bsd.port.mk>
