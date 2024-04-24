PORTNAME=	terraform-plugin-docs
PORTVERSION=	0.19.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Generate and validate Terraform plugin/provider documentation
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

GO_TARGET=	./cmd/tfplugindocs

PLIST_FILES=	bin/tfplugindocs

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
