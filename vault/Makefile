PORTNAME=	vault
PORTVERSION=	1.16.1
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	swills@FreeBSD.org
COMMENT=	Tool for securely accessing secrets
WWW=		https://vaultproject.io/

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		cpe go:modules
CPE_VENDOR=	hashicorp

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-tags=ui -ldflags="-s -w -extldflags '-static' \
		-X github.com/hashicorp/vault/sdk/version.Version=${PORTVERSION} \
		-X github.com/hashicorp/vault/sdk/version.VersionPrerelease= \
		-X github.com/hashicorp/vault/sdk/version.VersionMetadata= \
		-X github.com/hashicorp/vault/sdk/version.GitCommit='$$(${GIT_COMMIT})'"

USE_RC_SUBR=	vault

USERS=		vault
GROUPS=		vault

PLIST_FILES=	bin/${PORTNAME}
SUB_FILES=	pkg-message

PORTSCOUT=	limit:^(v?)[0-9]+\.[0-9]+\.[0-9]+$

post-patch: .PHONY
	@${CP} ${MASTERDIR}/files/http_bindata__assetfs.go ${WRKDIR}/vault-${PORTVERSION}/http/bindata_assetfs.go

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
