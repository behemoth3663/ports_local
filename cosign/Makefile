PORTNAME=	cosign
PORTVERSION=	3.1.3
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local security
PKGNAMESUFFIX=	-local

MAINTAINER=	bofh@FreeBSD.org
COMMENT=	Signing OCI containers and other artifacts using Sigstore
WWW=		https://www.sigstore.dev/

LICENSE=	APACHE20

USES=		cpe go:modules
CPE_VENDOR=	sigstore

GO_MODULE=	github.com/sigstore/cosign/v3
GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-buildid=
GO_LDFLAGS+=	-X sigs.k8s.io/release-utils/version.buildDate=${SOURCE_DATE_EPOCH_CMD:sh}
GO_LDFLAGS+=	-X sigs.k8s.io/release-utils/version.gitCommit=${COMMIT_CMD:sh:M*:[1]}
GO_LDFLAGS+=	-X sigs.k8s.io/release-utils/version.gitVersion=${PORTVERSION}
GO_LDFLAGS+=	-X sigs.k8s.io/release-utils/version.gitTreeState=clean
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry.go
	${RM} ${WRKSRC}/vendor/github.com/googleapis/gax-go/v2/telemetry.go

.include <bsd.port.pre.mk>

COMMIT_CMD=	git ls-remote https://${GO_MODULE:H} refs/tags/${DISTVERSIONFULL}^{}

# Bring DISTINFO_FILE into scope so we can get the timestamp.
SOURCE_DATE_EPOCH_CMD=	date -ur \
			$$(${GREP} -m1 TIMESTAMP ${DISTINFO_FILE} | ${CUT} -d' ' -f3) \
			'+%FT%TZ'

.include <bsd.port.post.mk>
