# 构建Stable Diffusion镜像
## 操作步骤
### 1. 拉取代码
```bash
cd ~
git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui.git
cd stable-diffusion-webui
git checkout 22bcc7be428c94e9408f589966c2040187245d81
```
### 2. 下载模型
```bash
cd ~
git clone https://huggingface.co/runwayml/stable-diffusion-v1-5
git lfs install
git lfs pull -include "v1-5-pruned-emaonly.safetensors"
# 将下载好的模型移动到Stable Diffusion项目的models下
mv majicmixRealistic_v5.safetensors ../stable-diffusion-webui/models/Stable-diffusion/
```
### 3. 构建镜像
```dockerfile
FROM nvidia/cuda:11.3.0-cudnn8-runtime-ubuntu20.04
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get install -y --no-install-recommends \
libgl1 libglib2.0-0 wget git curl vim python3 python3-venv && \
apt-get clean && \
rm -rf /var/lib/apt/lists/*

ADD . /stable-diffusion-webui
WORKDIR /stable-diffusion-webui/
RUN ./webui.sh -f can_run_as_root --exit --skip-torch-cuda-test

ENV VIRTUAL_ENV=/stable-diffusion-webui/venv
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

CMD ["python3", "launch.py", "--listen --skip-torch-cuda-test --no-half"]
```
```bash
docker build -t sd:v1 .
```
### 4. 本地测试
运行镜像
```bash
docker run -ti -p7860:7860 --rm --ipc=host sd:v1 bash
python3 launch.py --listen --skip-torch-cuda-test --no-half
```
在浏览器中访问 https://localhost:7860 