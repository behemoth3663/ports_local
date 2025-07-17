PORTNAME=	cloudsmith-api
PORTVERSION=    2.0.21
CATEGORIES=	local devel python
MASTER_SITES=	PYPI
PKGNAMEPREFIX=	${PYTHON_PKGNAMEPREFIX}
DISTNAME=	${PORTNAME:S/-/_/g}-${PORTVERSION}

MAINTAINER=	none@freebsd.org
COMMENT=	Cloudsmith API Clients (Generated)
WWW=		https://github.com/cloudsmith-io/cloudsmith-api

LICENSE=	APACHE20

RUN_DEPENDS=	${PYTHON_PKGNAMEPREFIX}click>=7.0:devel/py-click@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}certifi>=2017.4.17:security/py-certifi@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}python-dateutil>=2.1:devel/py-python-dateutil@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}six>=1.10:devel/py-six@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}urllib3>0:net/py-urllib3@${PY_FLAVOR}

USES=		python
USE_PYTHON=	autoplist concurrent pep517

.include <bsd.port.mk>
