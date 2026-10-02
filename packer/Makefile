PORTNAME=	packer
PORTVERSION=	1.16.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-local

MAINTAINER=	brad@facefault.org
COMMENT=	Tool for creating machine images
WWW=		https://packer.io https://${GO_MODULE}

LICENSE=	MPL20

USES=		go:modules

GO_MODULE=	github.com/hashicorp/packer

#CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-extldflags -static
GO_LDFLAGS+=	-X ${GO_MODULE}/version.GitCommit=${COMMIT_CMD:sh:M*:[1]}
GO_LDFLAGS+=	-X ${GO_MODULE}/version.Version=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	${RM} -r ${WRKSRC}/vendor/github.com/GoogleCloudPlatform
	${RM} -r ${WRKSRC}/vendor/go.opencensus.io 
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/experimental/opentelemetry
	${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/stats/opentelemetry
	${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry.go
	${RM} ${WRKSRC}/vendor/github.com/googleapis/gax-go/v2/telemetry.go

.include <bsd.port.pre.mk>

COMMIT_CMD=	git ls-remote https://${GO_MODULE} refs/tags/${DISTVERSIONFULL}

.include <bsd.port.post.mk>
