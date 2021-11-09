import {
    Row,
    Col,
    Card,
    Table,
    Button,
    Typography,
  } from "antd";
  import { useEffect, useMemo } from "react";
  import { useState } from "react/cjs/react.development";
  import api from "../../services/api";
  
  // Images
  const { Title } = Typography;

  // table code start
  const columns = [
    {
      title: "Vigência inicial",
      dataIndex: "vigini",
      key: "vigini",
      width: "32%",
    },
    {
      title: "Vigência final",
      dataIndex: "vigfim",
      key: "vigfim",
    },
  
    {
      title: "% Custo de Folha",
      key: "custofolha",
      dataIndex: "custofolha",
    },
    {
      title: "Ação",
      key: "action",
      dataIndex: "action",
    },
    
  ];

  function SimplesN() {
    const [tabelasAPI, setTabelasAPI] = useState([]);

    useEffect(() => {
      api.get("/tabelaSN")
        .then(response => setTabelasAPI(response.data.data))
    }, [])

    const dadosTabela = useMemo(() => {
      const dados = tabelasAPI.map((tabela,idx) => {
        return {
          key: tabela.id,
          vigini: (
            <>
              <div className="avatar-info">
                <Title level={5}>
                  {new Intl.DateTimeFormat('pt-BR').format(
                    new Date(tabela.data_inicial)
                  ) }
                </Title>
              </div>
            </>
          ),
          vigfim: (
            <>
              <div className="author-info">
                <Title level={5}>
                  {new Intl.DateTimeFormat('pt-BR').format(
                    new Date(tabela.data_final)
                  ) }
                </Title>
              </div>
            </>
          ),
      
          custofolha: (
            <>
              <div className="author-info">
              <Title level={5}>{tabela.target_folha} %</Title>
              </div>
            </>
          ),
          action: (
            <>
              <div className="author-info">
                <span>
                  <Button >Editar</Button>
                </span>
              </div>
            </>
          ),
        }
      });

      return dados;
    }, [tabelasAPI])

    return (
      <>
        <div className="tabled">
          <Row gutter={[24, 0]}>
            <Col xs="24" xl={24}>
              <Card
                bordered={false}
                className="criclebox tablespace mb-24"
                title="Tabelas Simples Nacional"
  
              >
                <div className="table-responsive">
                  <Table
                    columns={columns}
                    dataSource={dadosTabela}
                    pagination={false}
                    className="ant-border-space"
                  />
                </div>
              </Card>
  
            </Col>
          </Row>
        </div>
      </>
    );
  }
  
  export default SimplesN;