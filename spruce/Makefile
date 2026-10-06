PORTNAME=	spruce
DISTVERSIONPREFIX=	v
DISTVERSION=	1.35.21
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	YAML & JSON merging tool
WWW=		https://${GO_MODULE}

LICENSE=	MIT

USES=		go:modules

GO_MODULE=	github.com/geofffranks/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X main.Version=${DISTVERSIONFULL}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} ${WRKSRC}/vendor/github.com/googleapis/gax-go/v2/telemetry.go

.include <bsd.port.mk>
