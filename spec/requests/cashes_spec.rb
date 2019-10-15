# frozen_string_literal: true

RSpec.describe 'Cashes', type: :request do
  it_behaves_like 'new', uri: '/cashes/new.js'

  it_behaves_like 'destroy', uri: '/cashes/9.js' do
    let(:resource) { stub_model Cash }

    before { expect(Cash).to receive(:find).with('9').and_return(resource) }

    let(:success) { -> { should render_template :destroy } }
  end
end
