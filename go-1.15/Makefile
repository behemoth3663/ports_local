# Created by: Devon H. O'Dell <devon.odell@gmail.com>
# $FreeBSD$

PORTNAME=	go
PORTVERSION=	1.15.15
#PORTEPOCH=	1
CATEGORIES=	local lang
MASTER_SITES?=	https://golang.org/dl/ \
		https://github.com/dmgk/go-bootstrap/releases/download/${BOOTSTRAP_TAG}/:bootstrap \
		LOCAL/dmgk:bootstrap
PKGNAMESUFFIX=	-${PORTVERSION:R}
DISTFILES?=	go${PORTVERSION}.src.tar.gz \
		go-${OPSYS:tl}-${GOARCH_${ARCH}}${GOARM_${ARCH}}-${BOOTSTRAP_TAG}.tar.xz:bootstrap

MAINTAINER=	dmgk@FreeBSD.org
COMMENT=	Go programming language

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

.ifdef QEMU_EMULATING
IGNORE=		fails to build with qemu-user-static
.endif

ONLY_FOR_ARCHS=	aarch64 amd64 armv6 armv7 i386

RUN_DEPENDS=	${RUN_DEPENDS_${ARCH}}
# ld.bfd from devel/binutils is needed for working cgo on aarch64
RUN_DEPENDS_aarch64=	binutils>0:devel/binutils

TEST_DEPENDS=	${TEST_DEPENDS_${ARCH}}
TEST_DEPENDS_aarch64=	binutils>0:devel/binutils

USES=		shebangfix

SHEBANG_FILES=	misc/wasm/go_js_wasm_exec \
		src/net/http/cgi/testdata/test.cgi
SHEBANG_GLOB=	*.bash *.pl *.sh

CONFLICTS_INSTALL?=	go-devel

# Upstream archive contains files with UTF-8 names
EXTRACT_CMD=	${SETENV} LC_ALL=en_US.UTF-8 ${TAR}

.ifndef MASTERDIR
WRKSRC=		${WRKDIR}/go
.endif

OPTIONS_DEFINE_i386=	GO387

GO387_DESC=	Do not generate code with SSE2 (for old x86 CPU)

GO387_VARS=	GO386=387

BINARIES=	go gofmt
BOOTSTRAP_TAG=	go1.14

GOARCH_aarch64=	arm64
GOARCH_amd64=	amd64
GOARCH_armv6=	arm
GOARCH_armv7=	arm
GOARCH_i386=	386

GOARM_armv6=	6
GOARM_armv7=	7

.include <bsd.port.pre.mk>

.if ${ARCH} == aarch64 && ${OSVERSION} < 1200000
IGNORE=		aarch64 support requires FreeBSD 12.0 or newer
.endif

post-patch:
	@${REINPLACE_CMD} -e 's|^if ulimit -T|false \&\& &|' ${WRKSRC}/src/run.bash

pre-build:
	# Check that the running kernel has COMPAT_FREEBSD11 required by lang/go{,-devel} post-ino64
	@${SETENV} CC="${CC}" OPSYS="${OPSYS}" OSVERSION="${OSVERSION}" WRKDIR="${WRKDIR}" \
		${SH} ${SCRIPTSDIR}/rust-compat11-canary.sh

do-build:
	cd ${WRKSRC}/src ; ${SETENV} \
		XDG_CACHE_HOME=${WRKDIR} \
		GOROOT_BOOTSTRAP=${WRKDIR}/go-${OPSYS:tl}-${GOARCH_${ARCH}}${GOARM_${ARCH}}-bootstrap \
		GOROOT=${WRKSRC} \
		GOROOT_FINAL=${PREFIX}/${PORTNAME}${PKGNAMESUFFIX} \
		GOBIN= \
		GOOS=${OPSYS:tl} \
		GOARCH=${GOARCH_${ARCH}} \
		GO386=${GO386} \
		GOARM=${GOARM_${ARCH}} \
		${SH} make.bash -v

do-install:
	@cd ${WRKSRC} ; \
		${RM} -r .gitattributes .gitignore .github favicon.ico robots.txt \
		pkg/obj pkg/bootstrap pkg/${OPSYS:tl}_${GOARCH_${ARCH}}/cmd
	@${MKDIR} ${STAGEDIR}${PREFIX}/${PORTNAME}${PKGNAMESUFFIX}
	@${CP} -a ${WRKSRC}/* ${STAGEDIR}${PREFIX}/${PORTNAME}${PKGNAMESUFFIX}
	@cd ${STAGEDIR}${PREFIX} && ${FIND} ${PORTNAME}${PKGNAMESUFFIX} \
		-type f ! -name '*.orig' ! -name '*.bak' >> ${TMPPLIST}

do-test:
	cd ${WRKSRC}/src && ${SETENV} \
		GOROOT=${WRKSRC} \
		PATH=${WRKSRC}/bin:${PATH} \
		GOOS=${OPSYS:tl} \
		GOARCH=${GOARCH_${ARCH}} \
		GO386=${GO386} \
		GOARM=${GOARM_${ARCH}} \
		${SH} run.bash -no-rebuild

.include <bsd.port.post.mk>
