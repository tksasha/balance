# frozen_string_literal: true

RSpec.describe 'Categories', type: :request do
  it_behaves_like 'new', '/categories/new.js'

  it_behaves_like 'edit', '/categories/49/edit.js' do
    let(:resource) { stub_model Category }

    before { expect(Category).to receive(:find).with('49').and_return(resource) }
  end

  it_behaves_like 'create', '/categories.js' do
    let(:params) { { category: { name: '' } } }

    let(:resource_params) { acp(params).require(:category).permit! }

    let(:resource) { stub_model Category }

    before { expect(Category).to receive(:new).with(resource_params).and_return(resource) }

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end
end
