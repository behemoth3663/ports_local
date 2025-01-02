PORTNAME=	infracost
DISTVERSIONPREFIX=	v
DISTVERSION=	0.10.40
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	dutra@FreeBSD.org
COMMENT=	Cloud cost estimates for Terraform
WWW=		https://${GO_MODULE}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

ONLY_FOR_ARCHS=	aarch64 amd64 powerpc64 powerpc64le riscv64
ONLY_FOR_ARCHS_REASON=	Requires 64-bit processor

USES=		go:modules

GO_MODULE=	github.com/${PORTNAME}/${PORTNAME}
GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X ${GO_MODULE}/internal/version.Version=v${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

.include <bsd.port.mk>
