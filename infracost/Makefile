PORTNAME=	infracost
DISTVERSIONPREFIX=	v
DISTVERSION=	0.10.46
PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX=	-local

MAINTAINER=	dutra@FreeBSD.org
COMMENT=	Cloud cost estimates for Terraform
WWW=		https://${GO_MODULE}

LICENSE=	APACHE20

ONLY_FOR_ARCHS=	aarch64 amd64 powerpc64 powerpc64le riscv64
ONLY_FOR_ARCHS_REASON=	Requires 64-bit processor

USES=		go:modules

GO_MODULE=	github.com/${PORTNAME}/${PORTNAME}
GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X ${GO_MODULE}/internal/version.Version=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/internal/trace
	${RM} -r ${WRKSRC}/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/stats/opentelemetry
	${RM} ${WRKSRC}/vendor/google.golang.org/grpc/experimental/opentelemetry/trace_options.go
	${RM} ${WRKSRC}/vendor/github.com/hashicorp/go-getter/get_gcs.go

.include <bsd.port.mk>
