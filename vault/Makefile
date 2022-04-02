#

PORTNAME=	vault
PORTVERSION=	1.8.7
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	swills@FreeBSD.org
COMMENT=	Tool for securely accessing secrets

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=		cpe go:modules
CPE_VENDOR=	hashicorp

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

COMMIT_ID!=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS=	-tags=ui \
		-ldflags="-s -w \
		-X github.com/hashicorp/vault/sdk/version.Version=${PORTVERSION} \
		-X github.com/hashicorp/vault/sdk/version.VersionPrerelease= \
		-X github.com/hashicorp/vault/sdk/version.VersionMetadata= \
		-X github.com/hashicorp/vault/sdk/version.GitCommit='${COMMIT_ID:[1]}'"

USE_RC_SUBR=	vault

USERS=		vault
GROUPS=		vault

PLIST_FILES=	bin/${PORTNAME}
SUB_FILES=	pkg-message

PORTSCOUT=	limit:^(v?)[0-9]+\.[0-9]+\.[0-9]+$

post-patch: .PHONY
	@${CP} ${MASTERDIR}/files/http_bindata__assetfs.go ${WRKDIR}/vault-${PORTVERSION}/http/bindata_assetfs.go

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
