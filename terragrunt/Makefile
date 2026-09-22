PORTNAME=	terragrunt
DISTVERSIONPREFIX=	v
DISTVERSION=	1.1.6
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
GO_LDFLAGS+=	-w -X ${GO_MODULE}/internal/version.Version=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

.include <bsd.port.mk>
