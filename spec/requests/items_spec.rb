# frozen_string_literal: true

RSpec.describe 'Items', type: :request do
  it_behaves_like 'index', '/items'

  it_behaves_like 'index', '/items.js'

  it_behaves_like 'destroy', '/items/57.js' do
    let(:resource) { stub_model Item, date: '2019-10-15' }

    before { expect(Item).to receive(:find).with('57').and_return(resource) }

    let(:success) { -> { should render_template :destroy } }
  end

  it_behaves_like 'create', '/items.js' do
    let(:params) { { item: { date: '' } } }

    let(:resource_params) { acp(params).require(:item).permit! }

    let(:resource) { stub_model Item, date: '2019-10-15' }

    before { expect(Item).to receive(:new).with(resource_params).and_return(resource) }

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end
end
