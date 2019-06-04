# -*- coding: utf-8 -*-

from setuptools import setup
import os.path as path

with open(path.join('testindicator', 'VERSION')) as version_file:
	version = version_file.read().strip()

with open("README.rst", "rb") as f:
	long_descr = f.read().decode("utf-8")

setup(
	name='bake',
	packages=['bake'],
	include_package_data=True,
	zip_safe=False,
	install_requires=[
		'pyyaml'
	],
	entry_points={
		'console_scripts': [
			'bake = bake.bake:main',
		]
	},
	#data_files=[
	#	('share/applications', ['data/testindicator.desktop']),
	#],
	keywords='make bake recipes target',
	version=version,
	description="Don't worry about running tests - just write code",
	long_description=long_descr,
	author='Logi Leifsson',
	author_email='logileifs@gmail.com',
	url='https://github.com/logileifs/bake.git',
)
