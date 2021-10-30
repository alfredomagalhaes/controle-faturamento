/*!
=========================================================
* Muse Ant Design Dashboard - v1.0.0
=========================================================
* Product Page: https://www.creative-tim.com/product/muse-ant-design-dashboard
* Copyright 2021 Creative Tim (https://www.creative-tim.com)
* Licensed under MIT (https://github.com/creativetimofficial/muse-ant-design-dashboard/blob/main/LICENSE.md)
* Coded by Creative Tim
=========================================================
* The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
*/
import {
  Row,
  Col,
  Card,
  
  Table,
  
  message,
  
  Button,
  Avatar,
  Typography,
} from "antd";


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
  
];

const data = [
  {
    key: "1",
    vigini: (
      <>
        <div className="avatar-info">
          <Title level={5}>01/01/2017</Title>
        </div>
      </>
    ),
    vigfim: (
      <>
        <div className="author-info">
          <Title level={5}>31/12/2035</Title>
        </div>
      </>
    ),

    custofolha: (
      <>
        <div className="author-info">
          <p>29%</p>
          <span>
            <Button >Editar </Button>
          </span>
        </div>
      </>
    ),
   
  },
];


function SimplesN() {
  
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
                  dataSource={data}
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
