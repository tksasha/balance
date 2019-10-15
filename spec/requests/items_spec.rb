# frozen_string_literal: true

RSpec.describe 'Items', type: :request do
  let(:resource) { stub_model Item, date: '2019-10-15' }

  let :params do
    {
      item: {
        date: '',
        formula: '',
        category_id: '',
        description: ''
      }
    }
  end

  let(:resource_params) { acp(params).require(:item).permit! }

  it_behaves_like 'index', '/items'

  it_behaves_like 'index', '/items.js'

  it_behaves_like 'destroy', '/items/57.js' do
    before { expect(Item).to receive(:find).with('57').and_return(resource) }

    let(:success) { -> { should render_template :destroy } }
  end

  it_behaves_like 'create', '/items.js' do
    before { expect(Item).to receive(:new).with(resource_params).and_return(resource) }

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end

  it_behaves_like 'update', '/items/15.js' do
    before { expect(Item).to receive(:find).with('15').and_return(resource) }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :edit } }
  end
end
