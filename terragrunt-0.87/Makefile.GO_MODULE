PORTNAME=	terragrunt
DISTVERSIONPREFIX=	v
DISTVERSION=	0.87.7
#PORTREVISION=	1
CATEGORIES=	sysutils
PKGNAMESUFFIX?=	-${PORTVERSION:R}

MAINTAINER=	dutra@FreeBSD.org
COMMENT=	Terragrunt is a thin wrapper for Terraform
WWW=		https://terragrunt.gruntwork.io/

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

GO_MODULE=	github.com/gruntwork-io/terragrunt

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/gruntwork-io/go-commons/version.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

CONFLICTS=	sysutils/terragrunt

.include <bsd.port.mk>
