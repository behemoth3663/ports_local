PORTNAME=	terragrunt
PORTVERSION=	1.1.6
DISTVERSIONPREFIX=	v
PORTREVISION=	2
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

EXTRA_PATCHES=	${FILESDIR}/patch.diff:-p1

post-patch: .PHONY
	${RM} -r ${WRKSRC}/vendor/github.com/GoogleCloudPlatform
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} -r ${WRKSRC}/internal/telemetry/*_test.go
	${RM} -r ${WRKSRC}/internal/cas/probeorigin_test.go

.include <bsd.port.mk>
