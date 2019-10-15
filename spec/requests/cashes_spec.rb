# frozen_string_literal: true

RSpec.describe 'Cashes', type: :request do
  let(:resource) { stub_model Cash }

  let :params do
    {
      cash: {
        formula: '',
        name: ''
      }
    }
  end

  let(:resource_params) { acp(params).require(:cash).permit! }

  it_behaves_like 'new', '/cashes/new.js'

  it_behaves_like 'destroy', '/cashes/9.js' do
    before { expect(Cash).to receive(:find).with('9').and_return(resource) }

    let(:success) { -> { should render_template :destroy } }
  end

  it_behaves_like 'create', '/cashes.js' do
    before { expect(Cash).to receive(:new).with(resource_params).and_return(resource) }

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end

  it_behaves_like 'update', '/cashes/22.js' do
    before { expect(Cash).to receive(:find).with('22').and_return(resource) }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :edit } }
  end
end
