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
import * as moment from 'moment';
  
  // Images
  const { Title } = Typography;

  // table code start
  const columns = [
    {
      title: "Referência",
      dataIndex: "referencia",
      key: "referencia",
      width: "32%",
    },
    {
      title: "Valor faturado",
      dataIndex: "valfaturado",
      key: "valfaturado",
    },
  
    {
      title: "Data Lançamento",
      key: "dataisercao",
      dataIndex: "dataisercao",
    },
    {
      title: "Ação",
      key: "action",
      dataIndex: "action",
    },
    
  ];

  function Faturamento() {
    const [tabelasAPI, setTabelasAPI] = useState([]);

    useEffect(() => {
      api.get("/faturamento")
        .then(response => setTabelasAPI(response.data.data))
    }, [])

    const dadosTabela = useMemo(() => {
      const dados = tabelasAPI.map((tabela,idx) => {
        return {
          key: tabela.id,
          referencia: (
            <>
              <div className="avatar-info">
                <Title level={5}>
                    { moment( 
                            new Date(tabela.referencia.substring(0,4)
                                    +"-"
                                    +tabela.referencia.substring(tabela.referencia,4,6)
                                    +"-01T03:00:00"
                                    ) 
                            ).format("MM/YYYY")
                    }
                </Title>
              </div>
            </>
          ),
          valfaturado: (
            <>
              <div className="author-info">
                <Title level={5}>
                  {new Intl.NumberFormat('pt-BR',{
                          style: 'currency',
                          currency: "BRL"
                        }).format(tabela.valor_faturado)}
                </Title>
              </div>
            </>
          ),
      
          dataisercao: (
            <>
              <div className="author-info">
                <Title level={5}>
                  {new Intl.DateTimeFormat('pt-BR').format(
                    new Date(tabela.CreatedAt)
                  ) }
                </Title>
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
                title="Faturamentos mensais"
  
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
  
  export default Faturamento;