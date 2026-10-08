PORTNAME=	trufflehog
PORTVERSION=	3.99.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils
PKGNAMESUFFIX=	-local

MAINTAINER=	yuri@FreeBSD.org
COMMENT=	Find and verify credentials
WWW=		https://trufflesecurity.com/

LICENSE=	AGPLv3

USES=		go:modules

GO_MODULE=	github.com/trufflesecurity/${PORTNAME}/v3

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X ${GO_MODULE}/pkg/version.BuildVersion=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

post-extract: .PHONY
	${RM} -r ${WRKSRC}/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/stats/opentelemetry
	${RM} ${WRKSRC}/vendor/google.golang.org/grpc/experimental/opentelemetry/trace_options.go

.include <bsd.port.mk>
