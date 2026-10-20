PORTNAME=	terragrunt
PORTVERSION=	1.1.6
DISTVERSIONPREFIX=	v
PORTREVISION=	5
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-local

MAINTAINER=	dutra@FreeBSD.org
COMMENT=	Terragrunt is a thin wrapper for Terraform
WWW=		https://terragrunt.gruntwork.io/

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

BUILD_DEPENDS=	garble:local/garble

USES=		go:modules

GO_MODULE=	github.com/gruntwork-io/terragrunt

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X ${GO_MODULE}/internal/version.Version=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	${RM} -r ${WRKSRC}/internal/telemetry
	${RM} -r ${WRKSRC}/vendor/github.com/GoogleCloudPlatform
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/experimental/opentelemetry
	${RM} ${WRKSRC}/vendor/github.com/googleapis/gax-go/v2/telemetry.go
	${FIND} ${WRKSRC} -name '*_test.go' -type f -delete
	${MKDIR} ${WRKSRC}/internal/traceopts
	${CP} -r ${FILESDIR}/options.go ${WRKSRC}/internal/traceopts

do-build: .PHONY
	( cd ${GO_WRKSRC} && ${ECHO_MSG} "===>  Building ${PORTNAME} from ." && \
		${SETENVI} ${WRK_ENV} ${MAKE_ENV} ${GO_ENV} GOMAXPROCS=${MAKE_JOBS_NUMBER} GOPROXY=off \
			${LOCALBASE}/bin/garble -tiny build ${GO_BUILDFLAGS} -o ${GO_WRKDIR_BIN}/${PORTNAME} . )

.include <bsd.port.mk>
