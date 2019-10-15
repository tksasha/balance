# frozen_string_literal: true

RSpec.describe CashesController, type: :controller do
  describe '#collection' do
    before { expect(Cash).to receive(:order).with(:name).and_return(:collection) }

    its(:collection) { should eq :collection }
  end

  describe '#resource_params' do
    let(:params) { acp cash: { formula: nil, name: nil } }

    before { expect(subject).to receive(:params).and_return(params) }

    its(:resource_params) { should eq params[:cash].permit! }
  end

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { expect(subject).to receive(:params).and_return(id: 11) }

      before { expect(Cash).to receive(:find).with(11).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end

  it_behaves_like :update do
    let(:success) { -> { should render_template(:update).with_status(200) } }

    let(:failure) { -> { should render_template(:edit).with_status(422) } }
  end

  describe '#initialize_resource' do
    before { expect(Cash).to receive(:new).and_return(:resource) }

    before { subject.send :initialize_resource }

    its(:resource) { should eq :resource }
  end

  describe '#build_resource' do
    before { expect(subject).to receive(:resource_params).and_return(:resource_params) }

    before { expect(Cash).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end

  describe '#set_variant' do
    context do
      before { expect(subject).to receive(:params).and_return({}) }

      after { subject.send :set_variant }

      it { expect(subject).to_not receive(:request) }
    end

    context do
      before { expect(subject).to receive(:params).and_return(report: '') }

      after { subject.send :set_variant }

      it { expect(subject).to_not receive(:request) }
    end

    context do
      before { expect(subject).to receive(:params).and_return(report: '1') }

      after { subject.send :set_variant }

      it { expect(subject).to receive_message_chain('request.variant=').with(:report) }
    end
  end
end
