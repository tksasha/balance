# frozen_string_literal: true

RSpec.describe 'Items', type: :request do
  it_behaves_like 'index', uri: '/items'

  it_behaves_like 'index', uri: '/items.js'

  it_behaves_like 'destroy', uri: '/items/57.js' do
    let(:resource) { stub_model Item, date: '2019-10-15' }

    before { expect(Item).to receive(:find).with('57').and_return(resource) }

    let(:success) { -> { should render_template :destroy } }
  end
end
