require 'rails_helper'

RSpec.describe CashesController, type: :controller do
  it { should be_a ActsAsRESTController }

  describe '#collection' do
    before { expect(Cash).to receive(:order).with(:name).and_return(:collection) }

    its(:collection) { should eq :collection }
  end

  describe '#resource_params' do
    before { expect(subject).to receive(:params).and_return(acp cash: { formula: '', name: '' }) }

    its(:resource_params) { should eq permit! formula: '', name: '' }
  end

  it_behaves_like :update do
    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :edit } }
  end

  it_behaves_like :create do
    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end
end
