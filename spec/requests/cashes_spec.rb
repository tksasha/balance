# frozen_string_literal: true

RSpec.describe 'Cashes', type: :request do
  it_behaves_like 'new', '/cashes/new.js'

  it_behaves_like 'destroy', '/cashes/9.js' do
    let(:resource) { stub_model Cash }

    before { expect(Cash).to receive(:find).with('9').and_return(resource) }

    let(:success) { -> { should render_template :destroy } }
  end

  it_behaves_like 'create', '/cashes.js' do
    let :params do
      {
        cash: {
          formula: '',
          name: '',
        }
      }
    end

    let(:resource_params) { acp(params).require(:cash).permit! }

    let(:resource) { stub_model Cash }

    before { expect(Cash).to receive(:new).with(resource_params).and_return(resource) }

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end
end
