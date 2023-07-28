function getArticleContent (body) {
 
    let BLOCKS = 0;
 
    let CHANGE_RATE = 0.9;

    let MIN_LENGTH = 3;

    let html = body;



    let deleteLabel = function (html) {
        let regEx_script = /<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/g; // 定义script的正则表达式
        let regEx_style = /<style\b[^<]*(?:(?!<\/style>)<[^<]*)*<\/style>/g; // 定义style的正则表达式
        let regEx_html = /<(?:.|\s)*?>/g; // 定义HTML标签的正则表达式

        html = html.replace(regEx_script, "");
        html = html.replace(regEx_style, "");
        html = html.replace(regEx_html, "");
        html = html.replace("((\r\n)|\n)[\\s\t ]*(\\1)+", "$1").replace("^((\r\n)|\n)", "");//去除空白行
        html = html.replace("    +| +|　+", ""); //去除空白
        return html.trim();
    };

    let b_html = deleteLabel(html);

    let splitBlock = function (text) {
        let groupMap = new Array();
        let bais = text;
        let br = text.split('\n');
        let line = null,
            blocksLine = "";
        let theCount = 0,
            groupCount = 0,
            count = 0;//1.记录每次添加的行数；2.记录块号；3.记录总行数

        for (let i = 0; i < br.length; i++) {
            line = br[i];
            if (line != '') {
                if (line.length > MIN_LENGTH) {
                    if (theCount <= BLOCKS) {
                        blocksLine += line.trim();
                        theCount++;
                    }
                    else {
                        if (blocksLine != undefined) {
                            groupMap[groupCount] = blocksLine;
                            groupCount++;
                            blocksLine = line.trim();
                            theCount = 1;
                        }
                    }
                    count++;
                }
            }

        }

        if (theCount != 0 && blocksLine != undefined) {//加上没凑齐的给给定块数的
            groupMap[groupCount + 1] = blocksLine;
        }

        return groupMap;
    };

    let o_html = splitBlock(b_html);

    /**
     * 分析每块之间变化的情况
     * @param map 块集合
     * @return 正文
     */
    let judgeBlocks = function (map) {
        let sets = map;
        let contentBlock = [];
        let currentBlock = map.length; //当前行的长度
        let lastBlock = 0; //上一行的长度
        for (let i = 0; i < sets.length; i++) {
            if (sets[i] != undefined) {
                lastBlock = currentBlock;
                currentBlock = sets[i].length;
                let between = Math.abs(currentBlock - lastBlock) / Math.max(currentBlock, lastBlock);

                if (between >= CHANGE_RATE) {
                    contentBlock.push(i);
                }
            }
        }

        //下面是取多个峰值节点中两个节点之间内容长度最大的内容
        let matchNode = contentBlock.length;

        let lastContent = 0;//前一个两节点之间的内容长度
        let context = null;//结果
        if (matchNode > 2) {
            for (let i = 1; i < matchNode; i++) {
                let result = "";
                for (let j = contentBlock[i - 1]; j < contentBlock[i]; j++) {
                    result += map[j];
                }
                if (result.length > lastContent) {
                    lastContent = result.length;
                    context += result;
                }

            }
        }

        return context;
    };

    let articleContent = judgeBlocks(o_html);

    return articleContent;
}


function RedPillPush() {
    let content = getArticleContent(document.body.innerHTML);
    content = content.replace("nullLoading [MathJax]/extensions/MathMenu.js","")
    let url = window.location.href;
    const data = { ReadUrl: url, Content:content};

    fetch('http://127.0.0.1:8080/push_text?token=xxxxx', {
    method: 'POST',
    mode: 'no-cors',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => {
    console.log('Success:', data);
    })
    .catch((error) => {
    console.error('Error:', error);
    });
}

RedPillPush()