PORTNAME=	dumbproxy
PORTVERSION=	1.52.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Simple, scriptable, secure HTTP/SOCKS5 forward proxy
WWW=		https://${GO_MODULE}

LICENSE=	MIT

USES=		go:modules

GO_MODULE=	github.com/Snawoot/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X main.version=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

.include <bsd.port.mk>
