require 'rails_helper'

RSpec.describe CashesController, type: :controller do
  describe '#resource' do
    before { expect(subject).to receive(:params).and_return({ id: 31 }) }

    before { expect(Cash).to receive(:find).with(31).and_return(:resource) }

    its(:resource) { should eq :resource }
  end

  describe '#resource_params' do
    let(:params) do
      { cash: { name: 'name', formula: 123 } }
    end

    before { expect(subject).to receive(:params).and_return(acp params) }

    its(:resource_params) { should eq permit! params[:cash] }
  end

  it_behaves_like :new

  it_behaves_like :create do
    let(:resource) { stub_model Cash }

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :errors } }
  end

  it_behaves_like :edit

  it_behaves_like :update do
    let(:resource) { stub_model Cash }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :errors } }
  end

  it_behaves_like :destroy do
    let(:success) { -> { should render_template :destroy } }
  end
end
