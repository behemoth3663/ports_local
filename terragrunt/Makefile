PORTNAME=	terragrunt
PORTVERSION=	1.1.6
DISTVERSIONPREFIX=	v
PORTREVISION=	3
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-local

MAINTAINER=	dutra@FreeBSD.org
COMMENT=	Terragrunt is a thin wrapper for Terraform
WWW=		https://terragrunt.gruntwork.io/

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

GO_MODULE=	github.com/gruntwork-io/terragrunt

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X ${GO_MODULE}/internal/version.Version=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

pre-patch: .PHONY
	${RM} -r ${WRKSRC}/internal/telemetry
	${RM} -r ${WRKSRC}/vendor/github.com/GoogleCloudPlatform
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/experimental/opentelemetry
	${RM} ${WRKSRC}/vendor/github.com/googleapis/gax-go/v2/telemetry.go
	${FIND} ${WRKSRC} -name '*_test.go' -type f -delete
	${MKDIR} ${WRKSRC}/internal/traceopts
	${CP} -r ${FILESDIR}/options.go ${WRKSRC}/internal/traceopts

.include <bsd.port.mk>
