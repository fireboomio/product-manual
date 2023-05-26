# Docker部署

### Docker运行

<pre class="language-bash"><code class="lang-bash"><strong># 拉取镜像
</strong><strong>docker pull fireboomapi/fireboom_server:latest
</strong><strong># 运行镜像
</strong>docker run  -p 9123:9123 -p 9991:9991 -p 9992:9992 fireboomapi/fireboom_server:latest test
</code></pre>

打开控制面板，使用如下地址进行访问：

[http://localhost:9123](http://localhost:9123)
