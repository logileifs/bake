import sys
from os import getcwd
from os import listdir
from os import environ
from pathlib import Path
from subprocess import Popen
from subprocess import PIPE

import yaml


ALLOWED_NAMES = ['bakefile', 'recipes', '.bakefile', '.recipes']
ALLOWED_EXTENSIONS = ['.yml', '.yaml']


def get_kwargs(all_args):
	args = []
	kwargs = {}
	for arg in all_args:
		try:
			k, v = arg.split('=')
			kwargs[k] = v
		except ValueError:
			args.append(arg)
	return args, kwargs


def find_recipes_file(work_dir):
	files = listdir(work_dir)
	for file in files:
		p = Path(file)
		if p.stem in ALLOWED_NAMES:
			if p.suffix:
				if p.suffix in ALLOWED_EXTENSIONS:
					return p
			else:
				return p
	return None


def load_yaml(path):
	with open(str(path)) as stream:
		data = yaml.safe_load(stream)
	return data


def execute(command, work_dir):
	process = Popen(
		command,
		stdout=PIPE,
		stderr=PIPE,
		cwd=work_dir,
		env=environ.copy(),
		shell=True
	)
	process.wait()
	stdout, stderr = process.communicate()
	result = process.returncode
	if stdout:
		print(stdout.decode('utf-8'), end='')
	if stderr:
		print(stderr.decode('utf-8'), file=sys.stderr, end='')
	return result


def main():
	work_dir = getcwd()
	recipes_file = find_recipes_file(work_dir)
	data = load_yaml(recipes_file)
	recipe = sys.argv[1]
	all_args = sys.argv[2:]
	args, kwargs = get_kwargs(all_args)
	cmd = data[recipe]
	cmd_formatted = cmd.format(*args, **kwargs)
	result = execute(cmd_formatted, work_dir)
	return result


if __name__ == '__main__':
	result = main()
	sys.exit(result)
